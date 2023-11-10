import { AxiosError } from "axios";
import { toast } from "react-toastify";

export const errorHandler = (e:AxiosError) => {
  console.log('error',e)
    toast(JSON.stringify(e.response?.data) || e.message, {
        type: 'error',
      })
}