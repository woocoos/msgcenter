import { createStore } from 'ice';
import app from '@/models/app';
import user from '@/models/user';

export default createStore({
  app,
  user,
});
