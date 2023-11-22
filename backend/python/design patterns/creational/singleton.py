class SingleMeta(type):
    _instances = {}

    def __call__(self, *args, **kwargs):
        if self not in self._instances:
            instance = super().__call__(*args, **kwargs)
            self._instances[self] = instance

        return self._instances[self]


class Singleton(metaclass=SingleMeta):
    def some_business_logic(self):
        print("RUN")


s1 = Singleton()
s2 = Singleton()

print(id(s1), id(s2))
