import axios from 'axios';

const mainURL = "http://localhost:8001/admin"

const adminLogin = async (form) => {
    let admin = await axios.post(`${mainURL}/login`, form)
    return admin
}

export { adminLogin };