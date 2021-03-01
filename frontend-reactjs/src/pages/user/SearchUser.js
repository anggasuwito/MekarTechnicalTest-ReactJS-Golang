import React from 'react'

export default function SearchUser(props) {
    const { keywords, onKeywords } = props

    return (
        <div className="float-left">
            <div className="container">
                <input type="text" className="form-control" value={keywords} onChange={(e) => onKeywords(e.target.value)} placeholder="Search.... (Nama/No KTP)" />
            </div>
        </div>
    )
}