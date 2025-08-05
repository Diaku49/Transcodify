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

const sendResetPasswordEmail = (email) => {
    return api.post("/user/reset-password", {
        email
    })
}

const changePasswordByEmail = (token, password) => {
    return api.post("/user/change-password", {
        password
    }, {
        headers: {
            "Authorization": `Bearer ${token}`
        }
    })
}

export {
    Login,
    Signup,
    sendResetPasswordEmail,
    changePasswordByEmail
}