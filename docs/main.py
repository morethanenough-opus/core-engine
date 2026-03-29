from typing import List, Dict
from core_engine.exceptions import CoreEngineError

class CoreEngine:
    def __init__(self, config: Dict):
        self.config = config
        self.modules = {}

    def load_module(self, module_name: str, module_path: str) -> None:
        try:
            module = __import__(module_path)
            self.modules[module_name] = module
        except ImportError as e:
            raise CoreEngineError(f"Failed to load module {module_name}: {e}")

    def run(self) -> None:
        for module in self.modules.values():
            module.run(self.config)

    def shutdown(self) -> None:
        for module in self.modules.values():
            module.shutdown()

def main() -> None:
    config = {
        "modules": [
            {"name": "module1", "path": "core_engine.modules.module1"},
            {"name": "module2", "path": "core_engine.modules.module2"}
        ]
    }

    engine = CoreEngine(config)

    for module in config["modules"]:
        engine.load_module(module["name"], module["path"])

    try:
        engine.run()
    except KeyboardInterrupt:
        engine.shutdown()
        print("Shutdown complete.")

if __name__ == "__main__":
    main()