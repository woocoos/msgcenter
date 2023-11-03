import { Message } from "@/generated/msgsrv/graphql";
import { setItem } from "@/pkg/localStore";
import { createModel } from "ice";


type ModelState = {
  handshake: boolean;
  message: Message[];
};

export default createModel({
  state: {
    handshake: false,
    message: [] as Message[],
  },
  reducers: {
    setHandshake(prevState: ModelState, payload: boolean) {
      setItem('handshake', payload);
      prevState.handshake = payload;
    },
    setMessage(prevState: ModelState, payload: Message[]) {
      const data = payload ?? []
      setItem('message', data, 28800);
      prevState.message = data;
    },
  },
  effects: () => ({}),
});
