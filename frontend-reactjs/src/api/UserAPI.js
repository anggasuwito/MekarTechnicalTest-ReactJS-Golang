import axios from 'axios';

const mainURL = "http://localhost:8001/user"

const getAllUsers = async () => {
    let user = await axios.get(`${mainURL}/users`)
    return user
}

const getUserByID = async (id) => {
    let user = await axios.get(`${mainURL}/${id}`)
    return user
}

const createNewUser = async (form) => {
    let user = await axios.post(`${mainURL}/createuser`, form)
    return user
}

export { getAllUsers, getUserByID, createNewUser };