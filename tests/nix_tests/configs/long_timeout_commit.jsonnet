local default = import 'default.jsonnet';

default {
  'qie_1990-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '5s',
      },
    },
  },
}
