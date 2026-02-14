local config = import 'default.jsonnet';

config {
  'qie_1990-1'+: {
    config+: {
      storage: {
        discard_abci_responses: true,
      },
    },
  },
}
