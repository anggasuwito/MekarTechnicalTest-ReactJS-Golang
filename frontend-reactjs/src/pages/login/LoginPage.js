import React, { useState } from 'react'

const LoginPage = (props) => {
    const { onLogin, errorMessage } = props
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const handleUsername = (e) => {
        setUsername(e.target.value)
    }
    const handlePassword = (e) => {
        setPassword(e.target.value)
    }
    const handlesubmitButton = () => {
        onLogin(username, password)
    }
    let submitButton
    if (username !== "" && password !== "") {
        submitButton = <button className="btn btn-outline-primary" type="button" onClick={() => handlesubmitButton()} >Login</button>
    } else {
        submitButton = <button className="btn btn-outline-primary" type="button" onClick={() => handlesubmitButton()} disabled>Login</button>
    }
    let errorAlert
    if (errorMessage.length !== 0) {
        errorAlert = <div class="alert alert-danger" role="alert">
            {errorMessage}
        </div>
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
                        <input className="form-control" placeholder="Password" type="text" style={{ WebkitTextSecurity: "disc" }} value={password} onChange={handlePassword} />
                        <br />
                        <br />
                        {errorAlert}
                        <br /><br />
                        <div className="App">
                            {submitButton}
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