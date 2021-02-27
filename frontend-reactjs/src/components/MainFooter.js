import React, { Component } from 'react'

export default class MainFooter extends Component {
    render() {
        return (
            <div>
                <div className="container-fluid App-footer">
                    <footer>
                        <br />
                        <div className="d-flex justify-content-center">
                            <img src="https://image.flaticon.com/icons/svg/1383/1383259.svg" height="50" alt="" />&nbsp;&nbsp;
                        <img src="https://image.flaticon.com/icons/svg/1383/1383263.svg" height="50" alt="" />&nbsp;&nbsp;
                        <img src="https://image.flaticon.com/icons/svg/1383/1383269.svg" height="50" alt="" />&nbsp;&nbsp;
                        <img src="https://image.flaticon.com/icons/svg/2111/2111673.svg" height="50" alt="" />&nbsp;&nbsp;
                        <img src="https://image.flaticon.com/icons/svg/888/888903.svg" height="50" alt="" />&nbsp;&nbsp;
                        <img src="https://image.flaticon.com/icons/svg/1383/1383260.svg" height="50" alt="" />&nbsp;&nbsp;
                    </div>

                        <hr />
                        <div className="d-flex justify-content-center">
                            Copyright © 2021 Angga Q. Suwito
                    </div>
                        <br />
                    </footer>
                </div>
            </div>
        )
    }
}
