<?php

declare(strict_types=1);

namespace Backup8\Domain\ValueObject\Iam;

use JsonSerializable;
use DomainException;

/**
 * @OA\Schema(
 *   title="SecretAccessKey",
 *   required={"secretAccessKey"}
 * )
 */
class SecretAccessKey implements JsonSerializable
{
    /**
     * @OA\Property(type="string")
     */
    private string $secretAccessKey;

    public function __construct(string $secretAccessKey)
    {
        if (!$this->isValid($secretAccessKey)) {
            throw new DomainException('InvalidSecretAccessKey');
        }
        $this->secretAccessKey = $secretAccessKey;
    }

    private function isValid(string $value): bool
    {
        return preg_match(
            '/^(?<![A-Za-z0-9\/+=])[A-Za-z0-9\/+=]{40}(?![A-Za-z0-9\/+=])$/',
            $value
        ) === 1;
    }

    public function __toString()
    {
        return $this->secretAccessKey;
    }

    public function jsonSerialize(): string
    {
        return $this->secretAccessKey;
    }
}
