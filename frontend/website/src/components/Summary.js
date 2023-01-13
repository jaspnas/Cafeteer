import styles from "./../style.module.css";

const Summary = () => {
    return(
        <div className={`${styles.centerContent} ${styles.restaurantBoxContainer}`}>
            <div className={styles.restaurantBox}><span>Stuff here, box lookin' div1</span></div>
            <div className={styles.restaurantBox}><span>Stuff here, box lookin' div2</span></div>
            <div className={styles.restaurantBox}><span>Stuff here, box lookin' div3</span></div>
        </div>
    )
}

export default Summary
