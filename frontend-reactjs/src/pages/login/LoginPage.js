import React from 'react'

const LoginPage = (props) => {
    const { onLogin } = props
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
                        <input className="form-control" placeholder="Username" type="text" name="username" />
                        <br />
                        <input className="form-control" placeholder="Password" type="password" name="password" />
                        <br />
                        <br /><br />
                        <div className="App">
                            <button className="btn btn-outline-primary" type="button" onClick={onLogin}>Login</button>
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