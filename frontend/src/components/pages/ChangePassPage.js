import { useState, useEffect } from "react";
import { toast } from "react-toastify";
import { useNavigate, useSearchParams } from "react-router-dom";
import { changePasswordByEmail } from "../../services/userService";
import styles from "../../css/changepasspage.module.css";
import EyeSlashIcon from "../../assets/eye_slash.svg";
import EyeIcon from "../../assets/eye.png";

export default function ChangePassPage() {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [showPassword, setShowPassword] = useState(false);
    const [showConfirmPassword, setShowConfirmPassword] = useState(false);
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [token, setToken] = useState("");

    useEffect(() => {
        // Get token from URL parameters
        const tokenFromURL = searchParams.get("token");
        if (!tokenFromURL) {
            toast.error("Invalid or missing reset token");
            navigate("/auth");
            return;
        }
        setToken(tokenFromURL);
    }, [searchParams, navigate]);

    const handleSubmit = async (e) => {
        e.preventDefault();

        // Validation
        if (!password || !confirmPassword) {
            toast.error("Please fill in all fields");
            return;
        }

        if (password.length < 6) {
            toast.error("Password must be at least 6 characters long");
            return;
        }

        if (password !== confirmPassword) {
            toast.error("Passwords do not match");
            return;
        }

        setIsSubmitting(true);

        try {
            const response = await changePasswordByEmail(token, password);

            if (response.status === 202) {
                toast.success("Password changed successfully!");

                // Redirect to login page after successful password change
                setTimeout(() => {
                    navigate("/auth");
                }, 1500);
            } else {
                toast.error(response.data.message || "Failed to change password");
            }

        } catch (err) {
            if (err.response) {
                toast.error(err.response.data.message || "Failed to change password");
            } else if (err.request) {
                toast.error("Network error. Please check your connection.");
            } else {
                toast.error("An unexpected error occurred");
            }
        } finally {
            setIsSubmitting(false);
        }
    };

    const handleBackToLogin = () => {
        navigate("/auth");
    };

    return (
        <div className={styles.changePassContainer}>
            <div className={styles.changePassBox}>
                <h2>Change Your Password</h2>
                <p className={styles.description}>
                    Enter your new password below to complete the password reset process.
                </p>

                <form onSubmit={handleSubmit}>
                    <div className={styles.passwordInputWrapper}>
                        <label htmlFor="password">New Password</label>
                        <input
                            type={showPassword ? "text" : "password"}
                            id="password"
                            value={password}
                            onChange={e => setPassword(e.target.value)}
                            autoComplete="new-password"
                            required
                            placeholder="Enter your new password"
                            minLength={6}
                        />
                        <img
                            src={showPassword ? EyeSlashIcon : EyeIcon}
                            alt={showPassword ? "Hide password" : "Show password"}
                            className={styles.eyeIcon}
                            onClick={() => setShowPassword((prev) => !prev)}
                            style={{ cursor: "pointer" }}
                        />
                    </div>

                    <div className={styles.passwordInputWrapper}>
                        <label htmlFor="confirmPassword">Confirm New Password</label>
                        <input
                            type={showConfirmPassword ? "text" : "password"}
                            id="confirmPassword"
                            value={confirmPassword}
                            onChange={e => setConfirmPassword(e.target.value)}
                            autoComplete="new-password"
                            required
                            placeholder="Confirm your new password"
                            minLength={6}
                        />
                        <img
                            src={showConfirmPassword ? EyeSlashIcon : EyeIcon}
                            alt={showConfirmPassword ? "Hide password" : "Show password"}
                            className={styles.eyeIcon}
                            onClick={() => setShowConfirmPassword((prev) => !prev)}
                            style={{ cursor: "pointer" }}
                        />
                    </div>

                    <button
                        type="submit"
                        className={styles.changePassButton}
                        disabled={isSubmitting}
                    >
                        {isSubmitting ? "Changing Password..." : "Change Password"}
                    </button>
                </form>

                <div className={styles.backToLogin}>
                    <button
                        type="button"
                        className={styles.backButton}
                        onClick={handleBackToLogin}
                    >
                        Back to Login
                    </button>
                </div>
            </div>
        </div>
    );
}
