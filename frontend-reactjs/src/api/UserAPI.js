import axios from 'axios';

const mainURL = "http://localhost:8001/user"

const getAllUsers = async () => {
    const tokenKey = sessionStorage.getItem("auth")
    const headerConfig = {
        headers: {
            Authorization: `Bearer ${tokenKey}`
        }
    }
    let user = await axios.get(`${mainURL}/users`, headerConfig)
    return user
}

const getUserByID = async (id) => {
    const tokenKey = sessionStorage.getItem("auth")
    const headerConfig = {
        headers: {
            Authorization: `Bearer ${tokenKey}`
        }
    }
    let user = await axios.get(`${mainURL}/${id}`, headerConfig)
    return user
}

const createNewUser = async (form) => {
    const tokenKey = sessionStorage.getItem("auth")
    const headerConfig = {
        headers: {
            Authorization: `Bearer ${tokenKey}`
        }
    }
    let user = await axios.post(`${mainURL}/createuser`, form, headerConfig)
    return user
}

const updateUserByID = async (form, id) => {
    const tokenKey = sessionStorage.getItem("auth")
    const headerConfig = {
        headers: {
            Authorization: `Bearer ${tokenKey}`
        }
    }
    let user = await axios.post(`${mainURL}/updateuser/${id}`, form, headerConfig)
    return user
}

const deleteUserByID = async (id) => {
    const tokenKey = sessionStorage.getItem("auth")
    const headerConfig = {
        headers: {
            Authorization: `Bearer ${tokenKey}`
        }
    }
    let user = await axios.delete(`${mainURL}/deleteuser/${id}`, headerConfig)
    return user
}

export { getAllUsers, getUserByID, createNewUser, updateUserByID, deleteUserByID };