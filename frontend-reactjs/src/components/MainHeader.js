import React, { Component } from 'react'
import logo from '../assets/gif/logo.svg'

export default class MainHeader extends Component {
    render() {
        return (
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
            </header>
        )
    }
}
