import axios from 'axios';

const mainURL = "http://localhost:8001/study"

const getAllStudies = async () => {
    const tokenKey = sessionStorage.getItem("auth")
    const headerConfig = {
        headers: {
            Authorization: `Bearer ${tokenKey}`
        }
    }
    let study = await axios.get(`${mainURL}/studies`, headerConfig)
    return study
}

export { getAllStudies };