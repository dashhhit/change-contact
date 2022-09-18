import React, {useState} from "react";
import axios from "axios";
import '../style_form_file.css'
import { setAuthToken } from "../helpers/setAuthToken"

function Login() {
    var [state, setState] = useState("")
    let textWrongPass;
    textWrongPass = ""
    const handleSubmit = (token) => {
        //reqres registered sample user
        const loginPayload = {
            token: token,
        }

        var axios = require('axios');
        var data = JSON.stringify(loginPayload);

        var config = {
            method: 'post',
            url: 'http://127.0.0.1:8000/auth',
            headers: {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            data : data
        };

        axios(config)
            .then(function (response) {
                console.log(response.data, response)
                if (response.data.code === 200) {
                    const token = response.data.data;
                    console.log(response.data.data, "resp")

                    //set JWT token to local
                    localStorage.setItem("token", token);

                    //set token to axios common header
                    setAuthToken(token);

                    //redirect user to home page
                    window.location.href = '/'
                } else {
                    setState("Неверный токен")
                }

            })
            .catch(function (error) {
                console.log(error);
            });






        // console.log(token)
        // axios.post("http://127.0.0.1:8000/auth", loginPayload, config)
        //     .then(response => {
        //         //get token from response
        //
        //
        //     })
        //     .catch(err => console.log(err));
    };

    return (
        <form
            onSubmit={(event) => {
                event.preventDefault()
                console.log(event.target[0].value)
                const token =  event.target[0].value
                handleSubmit(token);
            }}
            className={"form_token"}
        >
            <label form="token" className={"content"}>Токен</label><br />
            <input type="text" id="token" name="token" className={"token-input"}/><br />
            <span>{state}</span> <br/>
            <input type="submit" value="Submit" className={'token_btn'}/>
        </form>
    );
}
export default Login