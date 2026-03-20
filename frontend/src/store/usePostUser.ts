import { create } from "zustand";
import type { PostUsuario, PostUsuarioDTO } from "../models/post-usuario.ts";
import {
  DELETE_UserPost,
  GET_UserPost,
  POST_UserPost,
  PUT_UserPost,
} from "../http/FetchingPostUser.ts";
import { useModalPost } from "./useModal.ts";
import { ObtenerToken } from "../helpers/JWTDecore.ts";
import { toast } from "sonner";

const initialPostUser = (): PostUsuario => ({
  post_user_id: 0,
  descripcion: "",
  usuario_id: 0,
  file: null,
  fileName: "",
  usuario_creacion: "",
  usuario_modificacion: "",
});

type Data = {
  form_post_user: PostUsuario;
  list_post_user: PostUsuarioDTO[];
  GetPostByUser: () => void;
  GetOne: (obj: PostUsuarioDTO) => void;
  SendData: () => void;
  DeletePost: (id: number) => void;
  isEditing: boolean;
  isLoading: boolean;
  isPDF: boolean;
  reset: () => void;
};

export const useUserPost = create<Data>()((set, get) => {
  return {
    form_post_user: initialPostUser(),
    list_post_user: [],
    isEditing: false,
    isLoading: false,
    isPDF: true,

    GetPostByUser: async () => {
      const result = await GET_UserPost(Number(ObtenerToken()?.user));

      if (result.success && Array.isArray(result.data?.resultado)) {
        set({
          list_post_user: result.data.resultado,
        });
      } else {
        set({
          list_post_user: [],
        });
      }

      return result.error || null;
    },

    GetOne: async (obj: PostUsuarioDTO) => {
      useModalPost.getState().openModal();
      set({
        form_post_user: {
          post_user_id: obj.post_user_id,
          descripcion: obj.descripcion,
          file: null,
          fileName: "",
          usuario_creacion: obj.usuario_creacion,
          usuario_modificacion: obj.usuario_modificacion,
          usuario_id: null,
        },
        isEditing: true,
      });
    },

    SendData: async () => {
      const { form_post_user, isEditing } = get();
      const token = ObtenerToken();

      if (!token) {
        toast("No se pudo obtener el token o decodificarlo.");
        return;
      }

      if (!isEditing && !form_post_user.file) {
        toast("Debe seleccionar un archivo PDF.");
        return;
      }

      if (form_post_user.file) {
        const isPDF = form_post_user.fileName.endsWith(".pdf");
        if (!isPDF) {
          toast("El archivo debe ser un PDF.");
          return;
        }
      }

      const formData = new FormData();
      formData.append("post_user_id", form_post_user.post_user_id.toString());
      formData.append("descripcion", form_post_user.descripcion);

      if (isEditing) {
        formData.append(
          "usuario_modificacion",
          form_post_user.usuario_modificacion,
        );
      } else {
        formData.append("usuario_creacion", form_post_user.usuario_creacion);
      }

      formData.append("usuario_id", token.user.toString());

      if (form_post_user.file) {
        const fileBlob = new Blob([form_post_user.file], {
          type: "application/pdf",
        });
        formData.append(
          "file",
          fileBlob,
          form_post_user.fileName || "file.pdf",
        );
      }

      let result;

      if (isEditing) {
        result = await PUT_UserPost(formData);
      } else {
        result = await POST_UserPost(formData);
      }

      if (isEditing) {
        useModalPost.getState().closeModal();
      }

      if (result.success) {
        toast(result.data.mensaje);
        get().reset();
        get().GetPostByUser();
      } else {
        get().reset();
        toast(result.error);
      }
    },

    DeletePost: async (id: number) => {
      const result = await DELETE_UserPost(id);

      if (result.success) {
        get().GetPostByUser();

        if (result.data.codigo === 200) {
          toast(result.data.mensaje);
        } else {
          toast(result.data.mensaje);
        }

        return result.data;
      }
    },

    reset: () => set({ form_post_user: initialPostUser(), isEditing: false }),
  };
});
