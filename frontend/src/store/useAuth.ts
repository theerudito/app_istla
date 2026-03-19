import { create } from "zustand";
import { POST_Login, POST_Register } from "../http/FetchingLogin.ts";
import type {LoginDTO, User} from "../models/usuario.ts";
import {toast} from "sonner";

const initialLogin = (): LoginDTO => ({
    identificacion: "",
    password: ""
});

const initialRegister = (): User => ({
    id_usuario: 0,
    identificacion: "",
    nombres: "",
    apellidos: "",
    email: "",
    password: "",
    id_perfil: 0
});

type Data = {
    form_login: LoginDTO;
    form_register: User;
    Login: () => void;
    Register: () => void;
    Logout: () => void;
    isLogin: boolean;
    reset: () => void;
};

export const useAuth = create<Data>((set, get) => ({
    form_login: initialLogin(),
    form_register: initialRegister(),
    isLogin: !!localStorage.getItem("token"),

    Login: async () => {
        const { form_login } = get();

        const result = await POST_Login(form_login);

        if (result.success) {
            localStorage.setItem("token", JSON.stringify(result.data?.mensaje));
            set({ isLogin: true });
            get().reset();
        }

        if (result.data.codigo != 200) {
            toast(result.data.mensaje);
            localStorage.removeItem("token");
            set({ isLogin: false });
            return;
        }
    },

    Register: async () => {
        const { form_register } = get();
        form_register.id_perfil = 1;

        const result = await POST_Register(form_register);

        if (result.success) {
            if (result.data.codigo === 200) {
                localStorage.setItem("token", JSON.stringify(result.data.mensaje));
                set({ isLogin: true });
                get().reset();
                return result.data;
            } else {
                toast(result.data.mensaje);
                localStorage.removeItem("token");
                set({ isLogin: false });
                return result.data;
            }
        } else {
            localStorage.removeItem("token");
            set({ isLogin: false });
            return;
        }
    },

    Logout: () => {
        localStorage.removeItem("token");
        set({ isLogin: false });
        get().reset();
    },

    reset: () => set({ form_login: initialLogin(), form_register: initialRegister() }),
}));