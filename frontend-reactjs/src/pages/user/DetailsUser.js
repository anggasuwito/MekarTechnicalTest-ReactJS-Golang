import React from 'react'
import { Modal, Button } from 'react-bootstrap'

const DetailsUser = (props) => {
    const { resultDetails, show, onHide } = props
    return (
        <Modal
            show={show}
            size="lg"
            aria-labelledby="contained-modal-title-vcenter"
            centered
        >
            <Modal.Header closeButton onClick={() => onHide()}>
                <Modal.Title id="contained-modal-title-vcenter">
                    Details User
            </Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <div>
                    <div className="container">
                        <br />
                        <div className="card ">

                            <div className="card-body">
                                <form>
                                    <br />
                                    <label>Nama</label>
                                    <input className="form-control" value={resultDetails.name} disabled type="text" />
                                    <br />
                                    <label>Tanggal Lahir</label>
                                    <input className="form-control" value={resultDetails.birthDate} disabled type="date" />
                                    <br />
                                    <label>No KTP</label>
                                    <input className="form-control" value={resultDetails.numberIdCard} disabled type="text" />
                                    <br />
                                    <label>Pekerjaan</label>
                                    <input className="form-control" value={resultDetails.work.name} disabled type="text" />
                                    <br />
                                    <label>Pendidikan Terakhir</label>
                                    <input className="form-control" value={resultDetails.study.name} disabled type="text" />
                                </form>
                                <br />
                            </div>
                        </div>
                    </div>
                </div>
            </Modal.Body>
            <Modal.Footer>
                <Button onClick={() => onHide()}>Close</Button>
            </Modal.Footer>
        </Modal>
    )
}

export default DetailsUser