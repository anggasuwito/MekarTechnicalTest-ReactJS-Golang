import React, { useState } from 'react'
import { Modal, Button } from 'react-bootstrap'

const UpdateUser = (props) => {
    const [id] = useState(props.id)
    const [name, setNama] = useState(props.name)
    const [tanggalLahir, setTanggalLahir] = useState(props.birthDate)
    const [noKTP, setNamaKTP] = useState(props.numberIdCard)
    const [pekerjaan, setPekerjaan] = useState(props.numberIdCard)
    const [pendidikanTerakhir, setPendidikanTerakhir] = useState(props.numberIdCard)

    const handleChangeInputNama = (event) => { setNama(event.target.value) }
    const handleChangeInputTanggalLahir = (event) => { setTanggalLahir(event.target.value) }
    const handleChangeInputNoKTP = (event) => { setNamaKTP(event.target.value) }
    const handleChangeInputPekerjaan = (event) => { setPekerjaan(event.target.value) }
    const handleChangeInputPendidikanTerakhir = (event) => { setPendidikanTerakhir(event.target.value) }

    const changeMenu = () => {
        props.updateMenuByID(name, tanggalLahir, noKTP, pekerjaan, pendidikanTerakhir, id)
        props.onHide()
    }

    const { show, onHide, works, studies } = props

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
                        </div>
                    </div>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="primary" onClick={() => changeMenu()}> Save</Button>
                    <Button variant="secondary" onClick={() => onHide()}>Cancel</Button>
                </Modal.Footer>
            </Modal>
        </div>
    )
}

export default UpdateUser