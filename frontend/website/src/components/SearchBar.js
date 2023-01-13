import styles from './../style.module.css'

const SearchBar = () => {
    return(
        <div className={`${styles.searchBarContainer} ${styles.centerContent}`}>
        <input className={styles.inputField}/>
        </div>
    )
}

export default SearchBar
