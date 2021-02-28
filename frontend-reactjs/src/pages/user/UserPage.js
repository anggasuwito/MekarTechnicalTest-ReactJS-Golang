import React, { useState, useEffect } from 'react'
import { getAllUsers, getUserByID, createNewUser } from '../../api/UserAPI'
import SearchUser from './SearchUser'
import UserPagination from './UserPagination'
import CreateUser from './CreateUser'
import UserList from './UserList'

const User = () => {
    const [keywords, setKeywords] = useState("")
    const [totalResult, setTotalResult] = useState(0)
    const [page, setPage] = useState(1)
    const [limit, setLimit] = useState(3)
    const [showCreateUserModal, setShowCreateUserModal] = useState(false)

    useEffect(() => {
        loadData()
    }, [])

    const loadData = () => {
        getAllUsers().then((response) => {
            console.log(response);
        })
    }

    const onKeywords = (e) => {
        setKeywords(e)
    }

    const onSetPage = (page) => {
        setPage(page)
    }

    const onSetLimit = (limit) => {
        setLimit(limit)
    }

    const handleCreateUserModal = () => {
        setShowCreateUserModal(!showCreateUserModal)
    }

    const addNewUser = (nama, tanggalLahir, noKTP, pekerjaan, pendidikanTerakhir) => {

    }
    return (
        <div className="container">
            <br /><br />
            <h1>TOTAL DATA = {totalResult}</h1><br /><br />
            <SearchUser keywords={keywords} onKeywords={onKeywords} />
            <UserPagination onSetLimit={onSetLimit} onSetPage={onSetPage} page={page} totalResult={totalResult} limit={limit} />
            <CreateUser show={showCreateUserModal} handleCreateUserModal={handleCreateUserModal} addNewUser={addNewUser} />
            <UserList />
            <br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br />
        </div>
    )
}

export default User
