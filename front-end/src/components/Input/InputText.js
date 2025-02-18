import React, {useState} from "react";
import "./InputText.css";

const InputText = (props) => {
    return (
        <>
            <label htmlFor={props.id}>
                {props.label}
            </label>
            <br/>
            <input
                id={props.id}
                type="text"
                placeholder={props.placeholder}
                className={props.className ? props.className : "appInput"}
                onChange={props.onChangeHandler}
            /> <br/>
        </>
    );
};

export default InputText;