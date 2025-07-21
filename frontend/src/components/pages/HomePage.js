import React from "react";
import styles from "../../css/homepage.module.css"

export default function HomePage() {


    return (
        <div className={styles.homepageContainer}>
            <h1>Welcome to Video Transcoder!</h1>
            <p>
                Effortlessly upload your videos and, well... enjoy!
            </p>
            <h3>Technologies Used</h3>
            <p>
                I used <b>RabbitMQ</b>, <b>GORM ORM</b>, <b>Redis</b>, and of course <b>FFmpeg</b> for transcoding videos.<br />
                I could have used WebSockets for live updates, but I refused to make my project even bigger—mainly because at this time I’m already in an internship and don’t have much time anyway (well, I do go outside on weekends, so that’s not totally true).<br /><br />
                Maybe I’ll add more features, maybe not—God knows.
            </p>
        </div>
    )
}