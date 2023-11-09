'use client'

import React ,{useState,FormEvent}from 'react';
import Auth from '../components/Auth';
import { authUrl } from '@/utils/network';
import axios, { AxiosError } from 'axios';
import { errorHandler } from '@/utils/errorHandler';
import { useRouter } from 'next/navigation';

type pageProps = {
    
};

const Login:React.FC<pageProps> = () => {
    const router = useRouter()
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
    
      const response = await axios.post(authUrl.login, arg)
    .catch((e:AxiosError)=>errorHandler(e))
     SetLoading(false);
  
  
    //   if (response.data) {
    //     localStorage.setItem(userTokenKey, response.data.token)
    //     router.push("/");
    //   }

    if (response) {
        
        router.push("/");
      }


    }
    
    return  <Auth
     loading={loading} 
     showRemembered
     title="Log In"
     buttonTitle = "Login"
     onSubmit={onSubmit} />;
}
export default Login;


