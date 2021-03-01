import axios from 'axios';

const mainURL = "http://localhost:8001/work"

const getAllWorks = async () => {
    const tokenKey = sessionStorage.getItem("auth")
    const headerConfig = {
        headers: {
            Authorization: `Bearer ${tokenKey}`
        }
    }
    let work = await axios.get(`${mainURL}/works`, headerConfig)
    return work
}

export { getAllWorks };