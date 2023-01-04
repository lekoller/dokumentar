import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:5000/api/v1",
});

export async function findAll() {
    const { data } = await api.get("/docs");
    await new Promise((resolve) => setTimeout(resolve, 2000));
    console.log(data);

    return data;
}

export async function write() {
    const { data } = await api.post("/docs", {});
    return data;
}
