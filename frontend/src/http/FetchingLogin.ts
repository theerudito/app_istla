import type {LoginDTO, User} from "../models/usuario.ts";
import axios from "axios";
import {url_base} from "./UrlBase.ts";
import type {ApiResponseAcciones} from "../models/ApiResponse.ts";


export const POST_Login = async (obj: LoginDTO) : Promise<{ success: true; data: ApiResponseAcciones }> => {
    try {
        const response = await axios.post<ApiResponseAcciones>(`${url_base}/login`, obj);

        return { success: true, data: response.data };

    } catch (error: unknown) {
        let mensajeError = "Error desconocido";
        if (axios.isAxiosError(error)) {
            mensajeError = error.response?.data?.mensaje || mensajeError;
        }
        return {
            success: true,
            data: {
                codigo: 500,
                mensaje: mensajeError,
            },
        };
    }
};

export const POST_Register = async (obj: User) : Promise<{ success: true; data: ApiResponseAcciones }> => {
    try {
        const response = await axios.post<ApiResponseAcciones>(`${url_base}/register`, obj);

        return { success: true, data: response.data };

    } catch (error: unknown) {
        let mensajeError = "Error desconocido";
        if (axios.isAxiosError(error)) {
            mensajeError = error.response?.data?.mensaje || mensajeError;
        }
        return {
            success: true,
            data: {
                codigo: 500,
                mensaje: mensajeError,
            },
        };
    }
};