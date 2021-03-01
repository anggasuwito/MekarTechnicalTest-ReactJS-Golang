import { Table } from "react-bootstrap/cjs";
import { Button } from 'react-bootstrap'
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTrash, faEye, faEdit } from "@fortawesome/free-solid-svg-icons";

const UserList = (props) => {
    const { result, showDetailsUser, deleteUser, showUpdateUser } = props
    return (
        <div style={{ marginTop: 70 }} className="container">
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>No</th>
                        <th>Nama</th>
                        <th>Tanggal Lahir</th>
                        <th>No KTP</th>
                        <th>Pekerjaan</th>
                        <th>Pendidikan Terakhir</th>
                        <th colSpan="3">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    {result.map((user, index) => {
                        return (<tr>
                            <td className="align-middle">{index + 1}</td>
                            <td className="align-middle">{user.name}</td>
                            <td className="align-middle">{user.birthDate}</td>
                            <td className="align-middle">{user.numberIdCard}</td>
                            <td className="align-middle">{user.work.name}</td>
                            <td className="align-middle">{user.study.name}</td>
                            <td>
                                <Button variant="secondary" onClick={() => { showDetailsUser(user.id) }} > <FontAwesomeIcon icon={faEye} /></Button>
                            </td>
                            <td>
                                <Button variant="success" onClick={() => { showUpdateUser(user) }}> <FontAwesomeIcon icon={faEdit} /></Button>
                            </td>
                            <td>
                                <Button variant="danger" onClick={() => { deleteUser(user.id) }}> <FontAwesomeIcon icon={faTrash} /></Button>
                            </td>
                        </tr>)
                    })
                    }
                </tbody>
            </Table>
        </div>
    )
}

export default UserList;
