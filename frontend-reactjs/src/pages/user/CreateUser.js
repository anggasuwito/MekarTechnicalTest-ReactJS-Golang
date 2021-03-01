import React, { useState } from 'react'
import { Modal, Button } from 'react-bootstrap'
const CreateUser = (props) => {
    const { show, handleCreateUserModal, addNewUser, works, studies } = props
    const [nama, setNama] = useState("")
    const [tanggalLahir, setTanggalLahir] = useState("")
    const [noKTP, setNamaKTP] = useState("")
    const [pekerjaan, setPekerjaan] = useState("")
    const [pendidikanTerakhir, setPendidikanTerakhir] = useState("")

    const handleChangeInputNama = (event) => { setNama(event.target.value) }
    const handleChangeInputTanggalLahir = (event) => { setTanggalLahir(event.target.value) }
    const handleChangeInputNoKTP = (event) => { setNamaKTP(event.target.value) }
    const handleChangeInputPekerjaan = (event) => { setPekerjaan(event.target.value) }
    const handleChangeInputPendidikanTerakhir = (event) => { setPendidikanTerakhir(event.target.value) }

    const submitNewUser = () => {
        addNewUser(nama, tanggalLahir, noKTP, pekerjaan, pendidikanTerakhir)
        closeCreateUser()
    }

    const closeCreateUser = () => {
        setNama("")
        setTanggalLahir("")
        setNamaKTP("")
        setPekerjaan("")
        setPendidikanTerakhir("")
        handleCreateUserModal()
    }
    let saveButton
    if (nama !== "" && tanggalLahir !== "" && noKTP !== "" && pekerjaan !== "" && pendidikanTerakhir !== "") {
        saveButton = <Button variant="primary" onClick={() => submitNewUser()} > Save</Button>
    } else {
        saveButton = <Button variant="primary" onClick={() => submitNewUser()} disabled> Save</Button>
    }
    return (
        <div className="float-right">
            <div className="container">
                <button className="btn btn-outline-primary" type="button" onClick={() => handleCreateUserModal()}>Create User</button>
            </div>
            <Modal
                show={show}
                size="lg"
                aria-labelledby="contained-modal-title-vcenter"
                centered
            >
                <Modal.Header closeButton onClick={() => closeCreateUser()}>
                    <Modal.Title id="contained-modal-title-vcenter">
                        Create New User
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <div>
                        <div className="container">
                            <br />
                            <div className="card ">

                                <div className="card-body">
                                    <form>
                                        <label>Nama *</label>
                                        <input className="form-control" onChange={handleChangeInputNama} value={nama} placeholder="Masukkan Nama" type="text" />
                                        <br />
                                        <label>Tanggal Lahir *</label>
                                        <input className="form-control" onChange={handleChangeInputTanggalLahir} value={tanggalLahir} placeholder="Masukkan Tanggal Lahir" type="date" />
                                        <br />
                                        <label>No KTP *</label>
                                        <input className="form-control" onChange={handleChangeInputNoKTP} value={noKTP} placeholder="Masukkan No KTP" type="text" />
                                        <br />
                                        <label>Pekerjaan *</label>
                                        <select class="form-control" onChange={handleChangeInputPekerjaan}>
                                            <option value={""}>--Pilih Pekerjaan--</option>
                                            {works.map((work) => {
                                                return (<option value={work.id}>{work.name}</option>)
                                            })}
                                        </select>
                                        <br />
                                        <label>Pendidikan Terakhir *</label>
                                        <select class="form-control" onChange={handleChangeInputPendidikanTerakhir}>
                                            <option value={""}>--Pilih Pendidikan Terakhir--</option>
                                            {studies.map((study) => {
                                                return (<option value={study.id}>{study.name}</option>)
                                            })}
                                        </select>
                                    </form>
                                    <br />
                                </div>
                            </div>
                            <br />  <br />  <br />
                        </div>
                    </div>
                </Modal.Body>
                <Modal.Footer>
                    {saveButton}
                    <Button variant="secondary" onClick={() => closeCreateUser()}>Cancel</Button>
                </Modal.Footer>
            </Modal>
        </div >
    )
}

export default CreateUser