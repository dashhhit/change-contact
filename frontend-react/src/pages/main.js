import '../style_form_token.css'
import React, {useState} from "react";

export default function Main() {

    const [selectedFile, setSelectedFile] = useState();
    const [isFilePicked, setIsFilePicked] = useState(false);

    const changeHandler = (event) => {
        setSelectedFile(event.target.files[0]);
        setIsFilePicked(true);
    };
    let linkRef;
    linkRef = React.createRef();
    var [state, setState] = useState("")
    var myHeaders = new Headers();
    var isValidToken = true
    var token = localStorage.getItem("token")
    myHeaders.append("Authorization", "Bearer " + token);

    const handleSubmit = () => {
        const formData = new FormData();
        formData.append('file', selectedFile)
        fetch('http://127.0.0.1:8000/getFile', {
            method: 'POST',
            mode: 'cors',
            credentials: 'same-origin',
            body: formData,
            headers: myHeaders,
            redirect: 'follow'
        })
            .then(res => {
                console.log(res)

                if (res.status === 200){
                    return res.blob()
                } else {
                    setState("Истекший токен")
                    localStorage.removeItem("token")
                    window.location.replace("http://127.0.0.1:3000/auth")

                }
        }).then( blob => {
            const href = window.URL.createObjectURL(blob);
            const a = linkRef.current;
            a.download = 'contacts.vcf';
            a.href = href;
            a.click();
            a.href = '';
        }).catch(err => console.error(err))
    };

    return (
        <div className={"form_file"}>
            <label className={"form_file_text"}>Выберите файл формата .vcf</label>
            <input type="file" accept=".vcf" onChange={changeHandler} className={"authentication"}/>
            <span>{state}</span> <br/>
            <div>
                <button onClick={handleSubmit} disabled={!isFilePicked} className={"file_btn"}>Отправить</button>
            </div>
            <a ref={linkRef}/>
            <a href={"http://127.0.0.1:3000/auth"} className={""}>Войти</a>
        </div>


    )
}

