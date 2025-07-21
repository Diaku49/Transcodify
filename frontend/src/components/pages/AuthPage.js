import { useState } from "react"
import styles from "../../css/authpage.module.css";
import { Login, Signup } from "../../services/userService";
import { useNavigate } from "react-router-dom";


export default function AuthPage() {
    const navigate = useNavigate();
    const [isLogin, setIsLogin] = useState(true);
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [status, setStatus] = useState(null);

    const handleSubmit = async (e) => {
        e.preventDefault()

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
                const responseData = await response.json()

                if (response.ok) {
                    setStatus(responseData.message)
                    const token = responseData.token
                    localStorage.setItem("token", token)

                    setTimeout(() => {
                        navigate("/")
                    }, 1500)
                } else {
                    setStatus(responseData.message)
                }
            }
            else {
                const response = await Signup(username, email, password, confirmPassword)
                const responseData = await response.json()
                if (response.ok) {
                    setStatus(responseData.message)
                    setIsLogin(true)
                } else {
                    setStatus(responseData.message)
                }
            }
        } catch (err) {
            console.log(err);
        }
    }

    return (<div className={styles.authContainer}>
        <div className={styles.authBox}>
            <h2>{isLogin ? "Login" : "Sign Up"}</h2>
            <form onSubmit={handleSubmit}>
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
                <div>
                    <label htmlFor="password">Password</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={e => setPassword(e.target.value)}
                        autoComplete={isLogin ? "current-password" : "new-password"}
                        required
                        placeholder="Your password"
                    />
                </div>
                {
                    !isLogin ? (<div>
                        <label htmlFor="confirmPassword">Confirm Password</label>
                        <input
                            type="password"
                            id="confirmPassword"
                            value={confirmPassword}
                            onChange={e => setConfirmPassword(e.target.value)}
                            required
                            placeholder="Confirm password"
                        />
                    </div>) : null
                }
                <button type="submit" className={styles.authButton}>
                    {isLogin ? "Login" : "Sign Up"}
                </button>
            </form>
            <p className={styles.switchText}>
                {isLogin ? "Don't have an account?" : "Already have an account?"}{" "}
                <button
                    type="button"
                    className={styles.switchButton}
                    onClick={() => setIsLogin((prev) => !prev)}
                >
                    {isLogin ? "Sign Up" : "Login"}
                </button>
            </p>
            {status && <p className={styles.status}>{status}</p>}
        </div>
    </div>)
}