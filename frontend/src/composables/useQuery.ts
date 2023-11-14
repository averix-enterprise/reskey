import { useQuery } from "@tanstack/vue-query";
import { GetAllHotKeys } from "../../wailsjs/go/backend/App";

export default () => {
    return {
        hotKeys: {
            all: () => useQuery({
                queryFn: () => GetAllHotKeys(),
                queryKey: ['hotKeys'],
            })
        }
    }
}