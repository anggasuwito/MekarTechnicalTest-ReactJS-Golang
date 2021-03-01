import React, { useState } from 'react'
import { Modal, Button } from 'react-bootstrap'

const UpdateUser = (props) => {
    const { result, show, onHide, works, studies, updateUser } = props
    const [id] = useState(result.id)
    const [name, setNama] = useState(result.name)
    const [tanggalLahir, setTanggalLahir] = useState(result.birthDate)
    const [noKTP, setNamaKTP] = useState(result.numberIdCard)
    const [pekerjaan, setPekerjaan] = useState(result.work.id)
    const [pendidikanTerakhir, setPendidikanTerakhir] = useState(result.study.id)

    const handleChangeInputNama = (event) => { setNama(event.target.value) }
    const handleChangeInputTanggalLahir = (event) => { setTanggalLahir(event.target.value) }
    const handleChangeInputNoKTP = (event) => { setNamaKTP(event.target.value) }
    const handleChangeInputPekerjaan = (event) => { setPekerjaan(event.target.value) }
    const handleChangeInputPendidikanTerakhir = (event) => { setPendidikanTerakhir(event.target.value) }

    const changeMenu = () => {
        updateUser(name, tanggalLahir, noKTP, pekerjaan, pendidikanTerakhir, id)
        onHide()
    }

    let saveButton
    if (name !== "" && tanggalLahir !== "" && noKTP !== "" && pekerjaan !== "" && pendidikanTerakhir !== "") {
        saveButton = <Button variant="primary" onClick={() => changeMenu()} > Save</Button>
    } else {
        saveButton = <Button variant="primary" onClick={() => changeMenu()} disabled> Save</Button>
    }
    return (
        <div>
            <Modal
                show={show}
                size="lg"
                aria-labelledby="contained-modal-title-vcenter"
                centered
            >
                <Modal.Header closeButton onClick={() => onHide()}>
                    <Modal.Title id="contained-modal-title-vcenter">
                        Update User
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
                                        <input className="form-control" onChange={handleChangeInputNama} value={name} placeholder="Masukkan Nama" type="text" />
                                        <br />
                                        <label>Tanggal Lahir *</label>
                                        <input className="form-control" onChange={handleChangeInputTanggalLahir} value={tanggalLahir} placeholder="Masukkan Tanggal Lahir" type="date" />
                                        <br />
                                        <label>No KTP *</label>
                                        <input className="form-control" onChange={handleChangeInputNoKTP} value={noKTP} placeholder="Masukkan No KTP" type="text" />
                                        <br />
                                        <label>Pekerjaan *</label>
                                        <select class="form-control" onChange={handleChangeInputPekerjaan} value={pekerjaan}>
                                            <option value={""}>--Pilih Pekerjaan--</option>
                                            {works.map((work) => {
                                                return (<option value={work.id}>{work.name}</option>)
                                            })}
                                        </select>
                                        <br />
                                        <label>Pendidikan Terakhir *</label>
                                        <select class="form-control" onChange={handleChangeInputPendidikanTerakhir} value={pendidikanTerakhir}>
                                            <option value={""}>--Pilih Pendidikan Terakhir--</option>
                                            {studies.map((study) => {
                                                return (<option value={study.id}>{study.name}</option>)
                                            })}
                                        </select>
                                    </form>
                                    <br />
                                </div>
                            </div>
                        </div>
                    </div>
                </Modal.Body>
                <Modal.Footer>
                    {saveButton}
                    <Button variant="secondary" onClick={() => onHide()}>Cancel</Button>
                </Modal.Footer>
            </Modal>
        </div>
    )
}

export default UpdateUser