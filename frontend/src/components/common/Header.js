import React from "react";
import userIcon from "../../assets/user.png";
import styles from "../../css/header.module.css";
import { Link } from "react-router-dom";

export default function Header() {
    const isLoggedIn = localStorage.getItem("token")
    const loginStatusColor = isLoggedIn ? "#22c55e" : "#ef4444";

    return (
        <header className={styles.header}>
            <div className={styles.headerLeft}></div>
            <div className={styles.headerCenter}>
                <Link className={styles.homeLink} to="/">Home</Link>
                <Link className={styles.uploadLink} to="/video/upload">Upload</Link>
                <Link className={styles.videoLink} to="/video">Videos</Link>
            </div>
            <div className={styles.headerRight}>
                <div className={styles.userIconWrapper}>
                    <Link className={styles.loginLink} to="/auth" title="Login / Signup">
                        <img src={userIcon} alt="Login / Signup" className={styles.userIcon} />
                    </Link>
                    <span
                        className={styles.userStatusIndicator}
                        style={{ backgroundColor: loginStatusColor }}
                        title={isLoggedIn ? "Logged In" : "Logged Out"}
                    ></span>
                </div>
            </div>
        </header>
    )
}