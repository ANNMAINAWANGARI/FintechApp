import { userTokenKey } from "@/utils/constants"
import { useRouter } from "next/navigation"
import { useContext } from "react"
import { ActionTypes, store } from "../components/StoreProvider"


const useLogout = () => {

    const Router = useRouter()
    const {dispatch} = useContext(store)

    const logout = () => {
        localStorage.removeItem(userTokenKey)
        dispatch({type:ActionTypes.UpdateUser, payload: null})
        Router.push("/login")
      }

    return {
        logout
    }
}

export default useLogout