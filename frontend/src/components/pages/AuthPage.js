import { useState } from "react";
import { toast } from "react-toastify";
import styles from "../../css/authpage.module.css";
import { Login, Signup } from "../../services/userService";
import { useNavigate } from "react-router-dom";
import EyeSlashIcon from "../../assets/eye_slash.svg";
import EyeIcon from "../../assets/eye.png";


export default function AuthPage({ setIsLoggedIn }) {
    const navigate = useNavigate();
    const [isLogin, setIsLogin] = useState(true);
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [isForgotPassword, setIsForgotPassword] = useState(false);
    const [showPassword, setShowPassword] = useState(false);
    const [showConfirmPassword, setShowConfirmPassword] = useState(false);

    const handleSubmit = async (e) => {
        e.preventDefault()

        if (isForgotPassword) {
            if (!email) {
                alert("Please enter your email.");
                return;
            }
            // Simulate API call for password reset and notif it with toastify
            setIsForgotPassword(false);
            return;
        }

        if (isLogin) {
            if (!email || !password) {
                alert("Please fill in all fields.");
                return;
            }
        } else {
            if (!email || !password || !username || !confirmPassword) {
                alert("Please fill in all fields.");
                return;
            }
        }

        try {
            if (isLogin) {
                const response = await Login(email, password)

                if (response.status === 202) {
                    toast.success(response.data.message)
                    const token = response.data.jwt
                    localStorage.setItem("token", token);
                    setIsLoggedIn(true)

                    setTimeout(() => {
                        navigate("/")
                    }, 2000)
                } else {
                    toast.error(response.data.message)
                }
            }
            else {
                const response = await Signup(username, email, password, confirmPassword)
                if (response.status === 201) {
                    toast.success(response.data.message)
                    setIsLogin(true)
                } else {
                    toast.error(response.data.message)
                }
            }
        } catch (err) {
            toast.error(err.response.data.message)
        }
    }

    return (<div className={styles.authContainer}>
        <div className={styles.authBox}>
            <h2>{isForgotPassword ? "Reset Password" : (isLogin ? "Login" : "Sign Up")}</h2>
            <form onSubmit={handleSubmit}>
                {isForgotPassword ? (
                    <>
                        <div>
                            <label htmlFor="email">Email</label>
                            <input
                                type="email"
                                id="email"
                                value={email}
                                onChange={e => setEmail(e.target.value)}
                                autoComplete="email"
                                required
                                placeholder="you@example.com"
                            />
                        </div>
                        <button type="submit" className={styles.authButton}>
                            Send Reset Link
                        </button>
                        <p className={styles.switchText}>
                            <button
                                type="button"
                                className={styles.switchButton}
                                // need to add the function for this
                                onClick={() => { setIsForgotPassword(false); }}
                            >
                                Back to Login
                            </button>
                        </p>
                    </>
                ) : (
                    <>
                        {!isLogin && (
                            <div>
                                <label htmlFor="username">Username</label>
                                <input
                                    type="text"
                                    id="username"
                                    value={username}
                                    onChange={e => setUsername(e.target.value)}
                                    autoComplete="username"
                                    required
                                    placeholder="Your username"
                                />
                            </div>
                        )}
                        <div>
                            <label htmlFor="email">Email</label>
                            <input
                                type="email"
                                id="email"
                                value={email}
                                onChange={e => setEmail(e.target.value)}
                                autoComplete="email"
                                required
                                placeholder="you@example.com"
                            />
                        </div>
                        <div className={styles.passwordInputWrapper}>
                            <label htmlFor="password">Password</label>
                            <input
                                type={showPassword ? "text" : "password"}
                                id="password"
                                value={password}
                                onChange={e => setPassword(e.target.value)}
                                autoComplete={isLogin ? "current-password" : "new-password"}
                                required
                                placeholder="Your password"
                            />
                            <img
                                src={showPassword ? EyeSlashIcon : EyeIcon}
                                alt={showPassword ? "Hide password" : "Show password"}
                                className={styles.eyeIcon}
                                onClick={() => setShowPassword((prev) => !prev)}
                                style={{ cursor: "pointer" }}
                            />
                        </div>
                        {
                            !isLogin ? (
                                <div className={styles.passwordInputWrapper}>
                                    <label htmlFor="confirmPassword">Confirm Password</label>
                                    <input
                                        type={showConfirmPassword ? "text" : "password"}
                                        id="confirmPassword"
                                        value={confirmPassword}
                                        onChange={e => setConfirmPassword(e.target.value)}
                                        required
                                        placeholder="Confirm password"
                                    />
                                    <img
                                        src={showConfirmPassword ? EyeSlashIcon : EyeIcon}
                                        alt={showConfirmPassword ? "Hide password" : "Show password"}
                                        className={styles.eyeIcon}
                                        onClick={() => setShowConfirmPassword((prev) => !prev)}
                                        style={{ cursor: "pointer" }}
                                    />
                                </div>
                            ) : null
                        }
                        <button type="submit" className={styles.authButton}>
                            {isLogin ? "Login" : "Sign Up"}
                        </button>
                        <p className={styles.switchText}>
                            {isLogin ? "Don't have an account?" : "Already have an account?"} {" "}
                            <button
                                type="button"
                                className={styles.switchButton}
                                onClick={() => { setIsLogin((prev) => !prev); }}
                            >
                                {isLogin ? "Sign Up" : "Login"}
                            </button>
                        </p>
                        {isLogin && (
                            <p className={styles.switchText}>
                                <button
                                    type="button"
                                    className={styles.switchButton}
                                    onClick={() => { setIsForgotPassword(true); }}
                                >
                                    Forgot Password?
                                </button>
                            </p>
                        )}
                    </>
                )}
            </form>
        </div>
    </div>)
}