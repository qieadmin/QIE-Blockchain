import pytest

from .network import setup_qie, setup_qie_rocksdb, setup_geth


@pytest.fixture(scope="session")
def qie(tmp_path_factory):
    path = tmp_path_factory.mktemp("qie")
    yield from setup_qie(path, 26650)


@pytest.fixture(scope="session")
def qie_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("qie-rocksdb")
    yield from setup_qie_rocksdb(path, 20650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["qie", "qie-ws"])
def qie_rpc_ws(request, qie):
    """
    run on both qie and qie websocket
    """
    provider = request.param
    if provider == "qie":
        yield qie
    elif provider == "qie-ws":
        qie_ws = qie.copy()
        qie_ws.use_websocket()
        yield qie_ws
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["qie", "qie-ws", "qie-rocksdb", "geth"])
def cluster(request, qie, qie_rocksdb, geth):
    """
    run on qie, qie websocket,
    qie built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "qie":
        yield qie
    elif provider == "qie-ws":
        qie_ws = qie.copy()
        qie_ws.use_websocket()
        yield qie_ws
    elif provider == "geth":
        yield geth
    elif provider == "qie-rocksdb":
        yield qie_rocksdb
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["qie", "qie-rocksdb"])
def qie_cluster(request, qie, qie_rocksdb):
    """
    run on qie default build &
    qie with rocksdb build and memIAVL + versionDB
    """
    provider = request.param
    if provider == "qie":
        yield qie
    elif provider == "qie-rocksdb":
        yield qie_rocksdb
    else:
        raise NotImplementedError
