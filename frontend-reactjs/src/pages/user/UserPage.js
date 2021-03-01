import React, { useState, useEffect } from 'react'
import { getAllUsers, getUserByID, createNewUser } from '../../api/UserAPI'
import { getAllStudies } from '../../api/StudyAPI'
import { getAllWorks } from '../../api/WorkAPI'
import Swal from 'sweetalert2'
import SearchUser from './SearchUser'
import UserPagination from './UserPagination'
import CreateUser from './CreateUser'
import DetailsUser from './DetailsUser'
import UserList from './UserList'

const User = (props) => {
    const { onLogout } = props
    const [keywords, setKeywords] = useState("")
    const [totalResult, setTotalResult] = useState(0)
    const [result, setResult] = useState([])
    const [resultDetails, setResultDetails] = useState({
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
    const [resultWorks, setResultWorks] = useState([])
    const [resultStudies, setResultStudies] = useState([])
    const [page, setPage] = useState(1)
    const [limit, setLimit] = useState(3)
    const [showCreateUserModal, setShowCreateUserModal] = useState(false)
    const [showDetailsUserModal, setShowDetailsUserModal] = useState(false)
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
    return (
        <div className="container">
            <br /><br />
            <h1>TOTAL DATA = {totalResult}</h1><br /><br />
            <SearchUser keywords={keywords} onKeywords={onKeywords} />
            <UserPagination onSetLimit={onSetLimit} onSetPage={onSetPage} page={page} totalResult={totalResult} limit={limit} />
            <CreateUser show={showCreateUserModal} handleCreateUserModal={handleCreateUserModal} addNewUser={addNewUser} works={resultWorks} studies={resultStudies} />
            <DetailsUser resultDetails={resultDetails} show={showDetailsUserModal} onHide={hideDetailsUser} />
            <UserList result={result} showDetailsUser={showDetailsUser} />
            <br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br />
        </div>
    )
}

export default User
