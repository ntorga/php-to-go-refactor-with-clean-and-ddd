<?php

declare(strict_types=1);

namespace Backup8\Domain\ValueObject\Iam;

use JsonSerializable;
use DomainException;

/**
 * @OA\Schema(
 *   title="AccessKeyId",
 *   required={"accessKeyId"}
 * )
 */
class AccessKeyId implements JsonSerializable
{
    /**
     * @OA\Property(type="string")
     */
    private string $accessKeyId;

    public function __construct(string $accessKeyId)
    {
        if (!$this->isValid($accessKeyId)) {
            throw new DomainException('InvalidAccessKeyId');
        }
        $this->accessKeyId = $accessKeyId;
    }

    private function isValid(string $value): bool
    {
        return preg_match('/^(?<![A-Z0-9])[A-Z0-9]{20}(?![A-Z0-9])$/', $value) === 1;
    }

    public function __toString()
    {
        return $this->accessKeyId;
    }

    public function jsonSerialize(): string
    {
        return $this->accessKeyId;
    }
}
