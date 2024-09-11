package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"superviseMe/core/entity"
	"superviseMe/core/module"
	"superviseMe/domain"

	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type userHandler struct {
	userUseCase module.UserUseCase
	oauthConfig *oauth2.Config
}

func NewUserHandler(userUseCase module.UserUseCase, oauthConfig *oauth2.Config) *userHandler {
	return &userHandler{userUseCase: userUseCase, oauthConfig: oauthConfig}
}

func (e *userHandler) LoginGoogle(w http.ResponseWriter, r *http.Request) {
	URL, err := url.Parse(domain.OAuthGoogleConf.Endpoint.AuthURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set required parameters
	parameters := url.Values{}
	parameters.Add("client_id", domain.OAuthGoogleConf.ClientID)
	parameters.Add("scope", strings.Join(domain.OAuthGoogleConf.Scopes, " "))
	parameters.Add("redirect_uri", domain.OAuthGoogleConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", domain.OAuthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *userHandler) getUserInfo(token *oauth2.Token) ([]byte, error) {
	client := h.oauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

func (e *userHandler) CallbackGoogle(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	log.Println("New Callback from Google OAuth :\n" + state)
	if state != domain.OAuthStateString {
		log.Printf("Invalid state. expected %s, got %s\n", domain.OAuthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		log.Println("Code not found")
		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("User permission has denied"))
			return
		}

		w.Write([]byte("Code Not Found to provide AccessToken"))
	} else {
		token, err := domain.OAuthGoogleConf.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Printf("OAuth Exchange failed : %v\n", err)
			return
		}

		log.Printf("[TOKEN_AUTH]Access Token : %s", token.AccessToken)
		log.Printf("[TOKEN_AUTH]Expiry Token : %s", token.Expiry.String())
		log.Printf("[TOKEN_AUTH]Refresh Token : %s", token.RefreshToken)

		userInfo, err := e.getUserInfo(token)
		if err != nil {
			http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var googleUser struct {
			ID      string `json:"id"`
			Email   string `json:"email"`
			Name    string `json:"name"`
			Picture string `json:"picture"`
		}

		err = json.Unmarshal(userInfo, &googleUser)
		if err != nil {
			http.Error(w, "Failed to parse user info: "+err.Error(), http.StatusInternalServerError)
			return
		}

		user := &entity.User{
			GoogleID: googleUser.ID,
			Email:    googleUser.Email,
			Name:     googleUser.Name,
			Picture:  googleUser.Picture,
		}

		err = e.userUseCase.SaveUser(user)
		if err != nil {
			http.Error(w, "Failed to save user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		claims := entity.MyClaims{
			Email:          user.Email,
			Id:             user.ID,
			StandardClaims: jwt.StandardClaims{},
		}
		tokenjwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, _ := tokenjwt.SignedString([]byte("Bolong"))

		respon := entity.ResponSucces{Message: "Succes", Token: signedToken, Data: user}
		result, err := json.Marshal(respon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			errResponse := entity.ResponsesError{Error: err.Error()}
			_ = json.NewEncoder(w).Encode(errResponse)
			return
		}

		fmt.Println("token :", tokenjwt)
		w.Write(result)
		fmt.Fprintf(w, "User Info: %+v\n", user)
	}
}

func (e *userHandler) Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user *entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = e.userUseCase.Registration(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func (e *userHandler) UpdateName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userGmail := r.Context().Value("email").(string)
	fmt.Println("ini", userGmail)
	var user *entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}
	err = e.userUseCase.UpdateName(userGmail, user.Name)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))

}

func (e *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user *entity.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user.Email)

	token, user, err := e.userUseCase.Login(user.Email, user.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.Write([]byte("Login Failed"))
		} else {
			w.Write([]byte(err.Error()))
		}
		return
	}
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	responsUser := entity.UserResponseToken{
		ID:      user.ID,
		Name:    user.Name,
		Picture: user.Picture,
		Email:   user.Email,
	}
	respon := entity.ResponSucces{Message: "Succes", Token: token, Data: responsUser}
	result, err := json.Marshal(respon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}

	fmt.Println("token :", token)
	w.Write(result)
}

func (e *userHandler) GetGoalSupervisor(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	userGmail := request.Context().Value("email").(string)
	user, err := e.userUseCase.GetGoalsBySuperviseeUser(userGmail)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
	}

	var supervisorGoalRespon []entity.GoalSupervisorRespons
	for _, s := range user.SupervisorGoals {
		supervisorGoalRespon = append(supervisorGoalRespon, entity.GoalSupervisorRespons{
			GoalName:        s.GoalName,
			SupervisorGmail: s.SupervisorGmail,
			CreatedAt:       s.CreatedAt,
			NilaiProgres:    s.NilaiProgres,
			GoalStatus:      s.GoalStatus,
		})
	}

	responsUser := entity.SupervisorRespons{
		Email:           user.Email,
		SupervisorGoals: supervisorGoalRespon,
	}

	responses := entity.ResponsesSucces{Message: "Succes", Data: responsUser}
	result, err := json.Marshal(responses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(result)

}

func (e *userHandler) GetGoalPersonal(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	userGmail := request.Context().Value("email").(string)
	user, err := e.userUseCase.GetGoalsBySuperviseeUser(userGmail)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
	}

	var personalGoalRespon []entity.GoalPersonalRespons
	for _, s := range user.PersonalGoals {
		personalGoalRespon = append(personalGoalRespon, entity.GoalPersonalRespons{
			GoalName:      s.GoalName,
			PersonalGmail: s.PersonalGmail,
			CreatedAt:     s.CreatedAt,
			NilaiProgres:  s.NilaiProgres,
			GoalStatus:    s.GoalStatus,
		})
	}

	responsUser := entity.PersonalRespons{
		Email:         user.Email,
		PersonalGoals: personalGoalRespon,
	}

	responses := entity.ResponsesSucces{Message: "Succes", Data: responsUser}
	result, err := json.Marshal(responses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(result)

}

func (e *userHandler) GetGoalsBySuperviseeUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	userGmail := request.Context().Value("email").(string)
	user, err := e.userUseCase.GetGoalsBySuperviseeUser(userGmail)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
	}
	var personalGoalsRespon []entity.GoalsResponseHome
	for _, i := range user.PersonalGoals {
		personalGoalsRespon = append(personalGoalsRespon, entity.GoalsResponseHome{
			ID:              i.ID,
			GoalName:        i.GoalName,
			Description:     i.Description,
			PersonalGmail:   i.PersonalGmail,
			SupervisorGmail: i.SupervisorGmail,
			CreatedAt:       i.CreatedAt,
			NilaiProgres:    i.NilaiProgres,
			GoalStatus:      i.GoalStatus,
		})
	}

	var supervisorGoalsRespon []entity.GoalsResponseHome
	for _, s := range user.SupervisorGoals {
		supervisorGoalsRespon = append(supervisorGoalsRespon, entity.GoalsResponseHome{
			ID:              s.ID,
			GoalName:        s.GoalName,
			Description:     s.Description,
			PersonalGmail:   s.PersonalGmail,
			SupervisorGmail: s.SupervisorGmail,
			CreatedAt:       s.CreatedAt,
			NilaiProgres:    s.NilaiProgres,
			GoalStatus:      s.GoalStatus,
		})
	}

	responsUser := entity.UserResponseHome{
		ID:              user.ID,
		Name:            user.Name,
		PersonalGoals:   personalGoalsRespon,
		SupervisorGoals: supervisorGoalsRespon,
	}

	responses := entity.ResponsesSucces{Message: "Succes", Data: responsUser}
	result, err := json.Marshal(responses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		errResponse := entity.ResponsesError{Error: err.Error()}
		_ = json.NewEncoder(writer).Encode(errResponse)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(result)

}
