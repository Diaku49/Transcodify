import { api } from "./api"

const Login = (email, password) => {
    return api.post("/user/login", {
        email,
        password
    })
}

const Signup = (userName, email, password, confirmPassword) => {
    return api.post("/user/signup", {
        userName,
        email,
        password,
        confirmPassword
    })
}

export {
    Login,
    Signup
}