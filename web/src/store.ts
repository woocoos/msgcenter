import { createStore } from 'ice';
import app from '@/models/app';
import user from '@/models/user';
import ws from '@/models/ws';

export default createStore({
  app,
  user,
  ws,
});
