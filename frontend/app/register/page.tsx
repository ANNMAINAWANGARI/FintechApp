'use client'

import React ,{useState,FormEvent}from 'react';
import Auth from '../components/Auth';
import { authUrl } from '@/utils/network';
import axios, { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import { useRouter } from 'next/navigation';
import { errorHandler } from '@/utils/errorHandler';
import withoutAuth from '../hocs/withoutAuth';

type pageProps = {
    
};

const Register:React.FC<pageProps> = () => {
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
    const response = await axios.post(authUrl.register, arg)
    .catch((e:AxiosError)=>errorHandler(e))
     SetLoading(false);
     
    if (response) {
      toast("User created successfully", {
        type: "success",
      });
      router.push("/login");
    }


    }
    
    return  <Auth
     loading={loading}
     title="Sign Up"
     buttonTitle = "Register"
     onSubmit={onSubmit}
     accountInfoText={{
        initialText:"Have an account",
        actionLink:"/login",
        actionText:"Login"
     }}
     />;
}
export default withoutAuth(Register);