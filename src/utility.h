#pragma once

#include <cstdlib>

namespace utility {
	template <class ActualClass>
	class Singleton
	{
	public:
		static ActualClass& Instance()
		{
			if (p == nullptr)
				p = new ActualClass;
			return *p;
		}
		Singleton(Singleton const &) = delete;
		void operator = (const Singleton&) = delete;
	protected:
		static ActualClass* p;
		Singleton() = default;
	};
	template <class T>
	T* Singleton<T>::p = nullptr;
}

int GetUniformRandom(int input_min, int input_max) {
	return input_min + rand() % (input_max - input_min);
}