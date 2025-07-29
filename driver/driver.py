from unittest.mock import MagicMock


class DBDriver:
    def execute(self, sql, vars):
        msg = self.prepare(sql, vars)
        self.conn.send(msg)
        data = self.conn.recv()
        return self.parse(data)

    def prepare(self, sql, vars):
        raise NotImplementedError()

    def parse(self, data):
        raise NotImplementedError()


class PGDriver(DBDriver):
    def prepare(self, sql, vars):
        print('type:', type(self))
        return b''  # FIXME

    def parse(self, data):
        return []  # FIXME


drv = PGDriver()
drv.conn = MagicMock()
drv.execute('SELECT name FROM users WHERE id=?', (900,))
