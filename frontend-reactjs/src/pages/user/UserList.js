import { Table } from "react-bootstrap/cjs";
import { Button } from 'react-bootstrap'
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTrash, faEye, faEdit } from "@fortawesome/free-solid-svg-icons";

const UserList = (props) => {
    return (
        <div style={{ marginTop: 70 }} className="container">
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>No.</th>
                        <th>Nama</th>
                        <th>Tanggal Lahir</th>
                        <th>No KTP</th>
                        <th>Pekerjaan</th>
                        <th>Pendidikan Terakhir</th>
                        <th colSpan="3">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        <tr>
                            <td>{"1"}</td>
                            <td>{"Nama"}</td>
                            <td>{"Tgl lahir"}</td>
                            <td>{"No KTP"}</td>
                            <td>{"Pekerjaan"}</td>
                            <td>{"Pendidikan Trakhir"}</td>
                            <td>
                                <Button variant="secondary" > <FontAwesomeIcon icon={faEye} /></Button>
                            </td>
                            <td>
                                <Button variant="success" > <FontAwesomeIcon icon={faEdit} /></Button>
                            </td>
                            <td>
                                <Button variant="danger" > <FontAwesomeIcon icon={faTrash} /></Button>
                            </td>
                        </tr>
                    }
                </tbody>
            </Table>
        </div>
    )
}

export default UserList;
