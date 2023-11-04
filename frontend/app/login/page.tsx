'use client'

import React ,{useState,FormEvent}from 'react';
import Auth from '../components/Auth';

type pageProps = {
    
};

const Login:React.FC<pageProps> = () => {
    const [loading, SetLoading] = useState(false);
    const onSubmit =async(
        e: FormEvent<HTMLFormElement>,
        formRef: React.RefObject<HTMLFormElement>
    )=>{
    e.preventDefault();
    SetLoading(true);
    let arg = {
      email: formRef.current?.email.value,
      password: formRef.current?.password.value,
    };


    }
    
    return  <Auth
     loading={loading} 
     showRemembered
     title="Log In"
     buttonTitle = "Login"
     onSubmit={onSubmit} />;
}
export default Login;