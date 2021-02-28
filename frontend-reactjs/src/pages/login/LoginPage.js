import React, { useState } from 'react'

const LoginPage = (props) => {
    const { onLogin } = props
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const handleUsername = (e) => {
        setUsername(e.target.value)
    }
    const handlePassword = (e) => {
        setPassword(e.target.value)
    }
    const submitButton = () => {
        onLogin(username,password)
    }
    return (
        <div className="container fluid col-4">
            <br /><br /><br />
            <div className="card">
                <div className="card-header">
                    <h3>Sign In</h3>
                </div>
                <div className="card-body">
                    <form>
                        <br />
                        <input className="form-control" placeholder="Username" type="text" value={username} onChange={handleUsername} />
                        <br />
                        <input className="form-control" placeholder="Password" type="password" value={password} onChange={handlePassword} />
                        <br />
                        <br /><br />
                        <div className="App">
                            <button className="btn btn-outline-primary" type="button" onClick={() => submitButton()}>Login</button>
                        </div>
                    </form>
                    <br /><br />
                </div>
            </div>
            <br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br />
        </div>
    )
}

export default LoginPage