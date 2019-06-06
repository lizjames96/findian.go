
package main
 
 int main()
{
	char str[] = "egalitarian";
	char search[] = "i";
	char *ptr = strstr(str, search);

	if (ptr != NULL)
	{
		printf("'%s' found '%s'\n", str, search);
	}
	else 
	{
		printf("'%s' not found '%s'\n", str, search);
	}

	return 0;
}
