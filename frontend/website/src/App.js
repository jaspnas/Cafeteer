import SearchBar from './components/SearchBar'
import DateField from './components/DateField'
import Summary from './components/Summary'
import styles from './style.module.css';


const App = () => {
    return(
        <div className={styles["centerContent"]}>
            <h1 className={styles["importantHeader"]}>Welcome!</h1>
            <SearchBar />
            <DateField />
            <Summary/>
        </div>
    )
}

export default App;
