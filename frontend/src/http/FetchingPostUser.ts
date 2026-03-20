import axios from "axios";
import { url_base } from "./UrlBase.ts";
import type { PostUsuarioDTO } from "../models/post-usuario.ts";
import type {
  ApiResponse,
  ApiResponseAcciones,
} from "../models/ApiResponse.ts";

export const GET_UserPost = async (id: number) => {
  try {
    const response = await axios.get<ApiResponse<PostUsuarioDTO[]>>(
      `${url_base}/post/get_by_user/${id}`,
    );
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};

export const POST_UserPost = async (
  obj: FormData,
): Promise<{ success: true; data: ApiResponseAcciones }> => {
  try {
    const response = await axios.post<ApiResponseAcciones>(
      `${url_base}/post`,
      obj,
      {
        headers: { "Content-Type": "multipart/form-data" },
      },
    );

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

export const PUT_UserPost = async (
  obj: FormData,
): Promise<{ success: true; data: ApiResponseAcciones }> => {
  try {
    const response = await axios.put<ApiResponseAcciones>(
      `${url_base}/post`,
      obj,
      {
        headers: { "Content-Type": "multipart/form-data" },
      },
    );

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

export const DELETE_UserPost = async (
  id: number,
): Promise<{ success: true; data: ApiResponseAcciones }> => {
  try {
    const response = await axios.delete<ApiResponseAcciones>(
      `${url_base}/post/${id}`,
    );

    return {
      success: true,
      data: response.data,
    };
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
