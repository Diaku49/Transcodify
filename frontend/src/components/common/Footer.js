import React from "react";
import styles from "../../css/footer.module.css";

export default function Footer() {
    const currentYear = new Date().getFullYear();

    return (
        <footer className={styles.footer}>
            © {currentYear} Video Transcoder &nbsp;|&nbsp;
            <a href="https://github.com/Diaku49" target="_blank" rel="noopener noreferrer">
                GitHub
            </a>
        </footer>
    )
}