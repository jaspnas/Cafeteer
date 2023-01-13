import styles from "./../style.module.css";

const DateField = () => {
    return(
        <div className={`${styles.searchBarContainer} ${styles.centerContent}`}>
            <input className={styles.inputField}/>
        </div>
    )
}

export default DateField
