import React, { useState, useEffect } from 'react'
import { getAllUsers, getUserByID, createNewUser, updateUserByID, deleteUserByID } from '../../api/UserAPI'
import { getAllStudies } from '../../api/StudyAPI'
import { getAllWorks } from '../../api/WorkAPI'
import Swal from 'sweetalert2'
import SearchUser from './SearchUser'
import UserPagination from './UserPagination'
import CreateUser from './CreateUser'
import DetailsUser from './DetailsUser'
import UserList from './UserList'
import UpdateUser from './UpdateUser'

const User = (props) => {
    const { onLogout } = props
    const [keywords, setKeywords] = useState("")
    const [totalResult, setTotalResult] = useState(0)
    const [result, setResult] = useState([])
    const [resultDetails, setResultDetails] = useState({
        id: 0,
        name: "",
        birthDate: "",
        numberIdCard: "",
        work: {
            id: 0,
            name: ""
        },
        study: {
            id: 0,
            name: ""
        }
    })
    const [resultWorks, setResultWorks] = useState([])
    const [resultStudies, setResultStudies] = useState([])
    const [page, setPage] = useState(1)
    const [limit, setLimit] = useState(3)
    const [showCreateUserModal, setShowCreateUserModal] = useState(false)
    const [showDetailsUserModal, setShowDetailsUserModal] = useState(false)
    const [showUpdateUserModal, setShowUpdateUserModal] = useState(false)

    useEffect(() => {
        loadData()
    }, [])// eslint-disable-line react-hooks/exhaustive-deps

    const loadData = () => {
        getAllUsers().then((response) => {
            if (response.data.meta.code === 202) {
                setTotalResult(response.data.meta.records)
                setResult(response.data.data)
            } else if (response.data.meta.code === 401) {
                expireTokenAlert()
            } else {
                errorAlert()
            }
        })
        getAllWorks().then((response) => {
            if (response.data.meta.code === 202) {
                setResultWorks(response.data.data)
            } else if (response.data.meta.code === 401) {
                expireTokenAlert()
            } else {
                errorAlert()
            }
        })
        getAllStudies().then((response) => {
            if (response.data.meta.code === 202) {
                setResultStudies(response.data.data)
            } else if (response.data.meta.code === 401) {
                expireTokenAlert()
            } else {
                errorAlert()
            }
        })
    }

    const successCreateNewUserAlert = () => {
        Swal.fire(
            'Success Create New User',
            '',
            'success'
        )
        loadData()
    }

    const successUpdateUserAlert = () => {
        Swal.fire(
            'Success Update User',
            '',
            'success'
        )
        loadData()
    }

    const successDeleteUserAlert = () => {
        Swal.fire(
            'Success Delete User',
            '',
            'success'
        )
        loadData()
    }

    const expireTokenAlert = () => {
        Swal.fire(
            'Your Session Has Expired !!!',
            'go to login page',
            'warning'
        ).then(() => {
            onLogout()
        })
    }

    const errorAlert = () => {
        Swal.fire(
            'Error',
            '',
            'error'
        )
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
        createNewUser({
            name: nama,
            birthDate: tanggalLahir,
            numberIdCard: noKTP,
            workId: Number(pekerjaan),
            studyId: Number(pendidikanTerakhir)
        }).then((response) => {
            if (response.data.meta.code === 202) {
                successCreateNewUserAlert()
            } else if (response.data.meta.code === 401) {
                expireTokenAlert()
            } else {
                errorAlert()
            }
        })
    }

    const updateUser = (nama, tanggalLahir, noKTP, pekerjaan, pendidikanTerakhir, id) => {
        updateUserByID({
            name: nama,
            birthDate: tanggalLahir,
            numberIdCard: noKTP,
            workId: Number(pekerjaan),
            studyId: Number(pendidikanTerakhir)
        }, id).then((response) => {
            if (response.data.meta.code === 202) {
                successUpdateUserAlert()
            } else if (response.data.meta.code === 401) {
                expireTokenAlert()
            } else {
                errorAlert()
            }
        })
    }

    const deleteUser = (id) => {
        Swal.fire({
            title: 'Are you sure want to delete this user ?',
            text: "You won't be able to revert this!",
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            confirmButtonText: 'Yes, delete it!'
        }).then((result) => {
            if (result.isConfirmed) {
                deleteUserByID(id).then((response) => {
                    if (response.data.meta.code === 202) {
                        successDeleteUserAlert()
                    } else if (response.data.meta.code === 401) {
                        expireTokenAlert()
                    } else {
                        errorAlert()
                    }
                })
            }
        })
    }

    const showDetailsUser = (id) => {
        getUserByID(id).then((response) => {
            if (response.data.meta.code === 202) {
                setResultDetails(response.data.data)
                setShowDetailsUserModal(!showDetailsUserModal)
            } else if (response.data.meta.code === 401) {
                expireTokenAlert()
            } else {
                errorAlert()
            }
        })
    }

    const hideDetailsUser = () => {
        setShowDetailsUserModal(!showDetailsUserModal)
        setResultDetails({
            id: 0,
            name: "",
            birthDate: "",
            numberIdCard: "",
            work: {
                id: "",
                name: ""
            },
            study: {
                id: "",
                name: ""
            }
        })
    }

    const showUpdateUser = (user) => {
        setResultDetails(user)
        setShowUpdateUserModal(!showUpdateUserModal)
    }

    const hideUpdateUser = () => {
        setShowUpdateUserModal(!showUpdateUserModal)
        setResultDetails({
            id: "",
            name: "",
            birthDate: "",
            numberIdCard: "",
            work: {
                id: "",
                name: ""
            },
            study: {
                id: "",
                name: ""
            }
        })
    }
    let updateModal
    if (showUpdateUserModal) {
        updateModal = <UpdateUser show={showUpdateUserModal} onHide={hideUpdateUser} result={resultDetails} updateUserByID={updateUser} works={resultWorks} studies={resultStudies} />
    }
    return (
        <div className="container">
            <br /><br />
            <h1>TOTAL DATA = {totalResult}</h1><br /><br />
            <SearchUser keywords={keywords} onKeywords={onKeywords} />
            <UserPagination onSetLimit={onSetLimit} onSetPage={onSetPage} page={page} totalResult={totalResult} limit={limit} />
            <CreateUser show={showCreateUserModal} handleCreateUserModal={handleCreateUserModal} addNewUser={addNewUser} works={resultWorks} studies={resultStudies} />
            <DetailsUser resultDetails={resultDetails} show={showDetailsUserModal} onHide={hideDetailsUser} />
            {updateModal}
            <UserList result={result} showDetailsUser={showDetailsUser} showUpdateUser={showUpdateUser} deleteUser={deleteUser} />
            <br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br />
        </div>
    )
}

export default User
