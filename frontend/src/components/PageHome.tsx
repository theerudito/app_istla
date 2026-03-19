import {Eye, Pencil, Trash2, Plus, Power} from "lucide-react";
import NewItemModal from "../modal/ModalPost.tsx";
import {useEffect, useState} from "react";
import {useUserPost} from "../store/usePostUser.ts";
import {useAuth} from "../store/useAuth.ts";
import {useModalPost} from "../store/useModal.ts";
import {ObtenerToken} from "../helpers/JWTDecore.ts";

export default function PagueHome() {
    const {Logout} = useAuth((state) => state);
    const {openModal} = useModalPost((state) => state);
    const {list_post_user, GetPostByUser, GetOne, DeletePost} = useUserPost((state) => state);
    const [showPDF, setShowPDF] = useState("");

    const nombre = ObtenerToken()

    useEffect(() => {
        GetPostByUser();
    }, [GetPostByUser])

    const handleView = (url: string) => {
       if (!url) return;
        setShowPDF(url);
    };

    return (
        <div className="flex flex-col h-screen bg-gray-900 text-gray-200">

            <div className="flex flex-col lg:flex-row flex-1 p-4 gap-4 overflow-hidden">

                <div className="w-full lg:w-1/2 bg-gray-800 rounded-xl shadow flex flex-col p-4">

                    <div className="mb-4">
                        <button
                            onClick={openModal}
                            className="flex items-center gap-2 bg-purple-500 px-4 py-2 rounded-lg hover:bg-purple-600 transition"
                        >
                            <Plus size={18}/>
                            Nuevo
                        </button>

                    </div>

                    <div
                        className="custom-scrollbar max-h-[45vh] lg:flex-1 lg:max-h-none overflow-y-auto border border-gray-700 rounded-lg">

                        <table className="w-full text-sm">

                            <thead className="bg-gray-700 sticky top-0 z-10">
                            <tr>
                                <th className="text-left p-3">ID</th>
                                <th className="text-left p-3">Descripción</th>
                                <th className="text-right p-3">Acciones</th>
                            </tr>
                            </thead>

                            <tbody>
                            {list_post_user.map((item) => (
                                <tr
                                    key={item.post_user_id}
                                    className="border-t border-gray-700 hover:bg-gray-700/40"
                                >
                                    <td className="p-3">{item.post_user_id}</td>

                                    <td className="p-3">
                                        {item.descripcion}
                                    </td>

                                    <td className="p-3">
                                        <div className="flex justify-end gap-3">

                                            <button
                                                onClick={() => handleView(item.url)}
                                                className="text-blue-400 hover:text-blue-300"
                                            >
                                                <Eye size={18}/>
                                            </button>

                                            <button
                                                onClick={() => GetOne(item)}
                                                className="text-yellow-400 hover:text-yellow-300"
                                            >
                                                <Pencil size={18}/>
                                            </button>

                                            <button
                                                onClick={() => DeletePost(item.post_user_id)}
                                                className="text-red-400 hover:text-red-300"
                                            >
                                                <Trash2 size={18}/>
                                            </button>

                                        </div>
                                    </td>

                                </tr>
                            ))}
                            </tbody>

                        </table>

                    </div>
                </div>

                <div
                    className="w-full lg:w-1/2 bg-gray-800 rounded-xl shadow border border-gray-700 overflow-hidden min-h-[300px]">

                    {showPDF && (
                        <iframe
                            src={showPDF}
                            className="w-full h-full bg-gray-900"
                        />
                    )}

                </div>

            </div>

            <footer className="bg-gray-800 border-t border-gray-700 py-3 px-4">
                <div className="flex items-center justify-between text-gray-200">

                    <span className="text-sm">
                        Bienvenido: <span className="text-purple-400 font-medium">{nombre?.nombres}</span>
                    </span>

                    <button
                        className="flex items-center gap-2 text-gray-300 hover:text-red-400 transition"
                        onClick={() => Logout()}
                    >
                        <Power size={18}/>
                        <span className="text-sm hidden sm:inline">Salir</span>
                    </button>

                </div>
            </footer>

            <NewItemModal/>

        </div>
    );
}